export default defineNuxtRouteMiddleware((to, from) => {
    // Allow public pages without auth
    if (to.meta.public) {
        return
    }
    // Redirect to login if not logged in
    const { loggedIn } = useUserSession()
    if (!loggedIn.value) {
        return navigateTo('/login')
    }
})
