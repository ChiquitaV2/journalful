export default defineNuxtRouteMiddleware(async (to, from) => {
    // Allow public pages without auth
    const { loggedIn, session, fetch } =  useUserSession()
    if (to.meta.public) {
        return
    }
    // Redirect to login if not logged in
    await fetch()
    console.log('loggedIn', loggedIn.value)
    console.log('session', session)
    if (!loggedIn.value) {
        return navigateTo('/login')
    }
})
