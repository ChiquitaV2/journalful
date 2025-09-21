// auth.d.ts
declare module '#auth-utils' {
    interface User {
        // Add your own fields
    }

    interface UserSession {
        // Add your own fields
    }

    interface SecureSessionData {
        accessToken: string
        refreshToken: string
        idToken: string
        expiresAt: number
        // Add your own fields
    }
}

export {}