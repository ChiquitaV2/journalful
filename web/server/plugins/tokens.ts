// /Users/owen/dev/projects/journalful/web/server/plugins/tokens.ts
import process from 'node:process'

import {
    sendError,
    getRequestURL,
    createError,
} from 'h3'
import type {NitroApp} from 'nitropack'
import type {H3Event} from 'h3'

// Define the structure of the token response from Zitadel
interface TokenResponse {
    access_token: string
    refresh_token?: string // Refresh token is sometimes optional on refresh
    expires_in: number
    token_type: string
}

// Define the structure of the introspection response from Zitadel
interface IntrospectionResponse {
    active: boolean
    scope: string
    client_id: string
    token_type: string
    exp: number
    iat: number
    sub: string
    aud: string[]
    iss: string
    jti: string
    username?: string
}

// --- Helper Functions ---

/**
 * Validates the access token against Zitadel's introspection endpoint.
 * @returns `true` if the token is active, `false` otherwise.
 */
async function validateAccessToken(
    accessToken: string,
    introspectionUrl: string,
    credentials: string,
): Promise<boolean> {
    try {
        const introspectionBody = new URLSearchParams({token: accessToken})
        const response = await $fetch<IntrospectionResponse>(introspectionUrl, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
                'Authorization': `Basic ${credentials}`,
            },
            body: introspectionBody.toString(),
        })
        return response.active
    } catch (error) {
        console.error('Error during token introspection:', error)
        // If introspection fails, it's safer to assume the token is invalid/expired.
        return false
    }
}

/**
 * Refreshes the access token using the refresh token.
 * @returns The new token response, or `null` if refreshing fails.
 */
async function refreshSessionTokens(
    refreshToken: string,
    tokenUrl: string,
    clientId: string,
    clientSecret: string,
): Promise<TokenResponse | null> {
    try {
        console.log('Access token expired. Attempting to refresh...')
        const refreshBody = new URLSearchParams({
            grant_type: 'refresh_token',
            refresh_token: refreshToken,
            client_id: clientId,
            client_secret: clientSecret,
        })

        const newTokens = await $fetch<TokenResponse>(tokenUrl, {
            method: 'POST',
            headers: {'Content-Type': 'application/x-www-form-urlencoded'},
            body: refreshBody.toString(),
        })

        console.log('Token refreshed successfully.')
        return newTokens
    } catch (error) {
        console.error('Failed to refresh token:', error)
        // This likely means the refresh token is also invalid.
        return null
    }
}

/**
 * Updates the user's session cookie with the new tokens.
 */
async function updateUserSession(
    event: H3Event,
    newTokens: TokenResponse,
    oldRefreshToken: string,
) {
    const expiresAt = Date.now() + newTokens.expires_in * 1000
    await setUserSession(event, {
        secure: {
            accessToken: newTokens.access_token,
            // Use the new refresh token if provided, otherwise keep the old one.
            refreshToken: newTokens.refresh_token || oldRefreshToken,
            expiresAt,
        },
    })
}

// --- Main Plugin Logic ---

export default defineNitroPlugin((nitroApp: NitroApp) => {
    nitroApp.hooks.hook('request', async (event) => {
        const url = getRequestURL(event)

        // We only want to run this logic for our own API routes.
        // Exclude auth routes to prevent an infinite loop during login/logout.
        if (!url.pathname.startsWith('/api/') || url.pathname.startsWith('/api/auth/'))
            return

        // --- 1. Get Config and Session ---
        const {
            NUXT_OAUTH_ZITADEL_DOMAIN,
            NUXT_OAUTH_ZITADEL_CLIENT_ID,
            NUXT_OAUTH_ZITADEL_CLIENT_SECRET,
        } = process.env

        if (!NUXT_OAUTH_ZITADEL_DOMAIN || !NUXT_OAUTH_ZITADEL_CLIENT_ID || !NUXT_OAUTH_ZITADEL_CLIENT_SECRET) {
            console.error('Zitadel OAuth environment variables are not configured.')
            return sendError(event, createError({statusCode: 500, statusMessage: 'Server configuration error.'}))
        }

        const session = await getUserSession(event)
        const accessToken = session.secure?.accessToken
        const refreshToken = session.secure?.refreshToken

        // If there's no access token, the user is not logged in.
        if (!accessToken)
            return

        const introspectionUrl = `https://${NUXT_OAUTH_ZITADEL_DOMAIN}/oauth/v2/introspect`
        const tokenUrl = `https://${NUXT_OAUTH_ZITADEL_DOMAIN}/oauth/v2/token`
        const credentials = Buffer.from(`${NUXT_OAUTH_ZITADEL_CLIENT_ID}:${NUXT_OAUTH_ZITADEL_CLIENT_SECRET}`).toString('base64')

        // --- 2. Validate Token ---
        const isTokenValid = await validateAccessToken(accessToken, introspectionUrl, credentials)
        if (isTokenValid)
            return // Token is valid, proceed with the request.

        // --- 3. Refresh Token ---
        // If we reach here, the access token is invalid or expired.
        if (!refreshToken) {
            await clearUserSession(event) // Clear any partial session data.
            return sendError(event, createError({statusCode: 401, statusMessage: 'Unauthorized: Session expired.'}))
        }

        const newTokens = await refreshSessionTokens(refreshToken, tokenUrl, NUXT_OAUTH_ZITADEL_CLIENT_ID, NUXT_OAUTH_ZITADEL_CLIENT_SECRET)

        if (!newTokens) {
            // Refresh failed, session is unrecoverable.
            await clearUserSession(event)
            return sendError(event, createError({statusCode: 401, statusMessage: 'Unauthorized: Invalid session.'}))
        }

        // --- 4. Update Session ---
        await updateUserSession(event, newTokens, refreshToken)
    })
})