export default defineOAuthZitadelEventHandler({
    config: {
        scope: ['openid', 'email', 'profile'],
    },
    async onSuccess(event, {user, tokens}) {
        console.log('User: ', user)
        console.log('Tokens: ', tokens)

        await setUserSession(event, {
            user: user,
            secure: {
                refreshToken: tokens.refresh_token,
                accessToken: tokens.access_token,
                idToken: tokens.id_token
            }
        });
        return sendRedirect(event, "/")
    },
    onError(event, error) {
        console.error('OAuth error:', error)
        return sendRedirect(event, '/')
    },

})