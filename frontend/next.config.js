/** @type {import('next').NextConfig} */
const nextConfig = {
    experimental: {
        serverActions: {
          // edit: updated to new key. Was previously `allowedForwardedHosts`
          allowedOrigins: ['localhost'],
        },
      },
}

module.exports = nextConfig
