export default defineNuxtRouteMiddleware(async (to, from) => {
  const { loggedIn, user } = useUserSession()
  
  // Only run for authenticated users
  if (!loggedIn.value || !user.value) {
    return
  }
  
  // Skip setup check for certain routes
  const skipRoutes = ['/login', '/setup', '/api', '/_', '/auth']
  if (skipRoutes.some(route => to.path.startsWith(route))) {
    return
  }
  
  try {
    // Use the dedicated setup status endpoint
    const setupStatus = await $fetch('/api/setup/status', {
      headers: useRequestHeaders(['cookie'])
    })
    
    // If user needs setup, redirect to setup page
    if (setupStatus.needsSetup) {
      console.log('User needs setup, redirecting...')
      return navigateTo('/setup')
    }
  } catch (error: any) {
    // For errors, let the user continue to avoid redirect loops
    console.warn('Setup status check failed:', error)
    return
  }
})