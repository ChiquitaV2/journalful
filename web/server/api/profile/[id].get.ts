import {generateMockProfile} from "~~/server/api/mock/data";

export default defineEventHandler(event => {
  const id = Number(event.context.params?.id) || 1

  return generateMockProfile(id, {
    name: 'Dr. Alex Johnson',
    bio: 'Machine Learning Researcher focused on academic literature analysis and citation networks. Passionate about developing tools that help researchers discover and organize scientific knowledge.',
    institution: 'University of Research Excellence',
    email: 'alex.johnson@university.edu',
    orcid: '0000-0000-0000-0000',
    website: 'https://alexjohnson.academic',
    location: 'San Francisco, CA',
    joinedAt: '2024-01-01'
  })
})