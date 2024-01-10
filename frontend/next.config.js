/** @type {import('next').NextConfig} */
const nextConfig = {
    experimental: {
        serverActions: {
          // edit: updated to new key. Was previously `allowedForwardedHosts`
          allowedOrigins: ['192.168.0.15:7777'],
        },
      },
}

module.exports = nextConfig
