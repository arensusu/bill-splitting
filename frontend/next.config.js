/** @type {import('next').NextConfig} */
const nextConfig = {
    experimental: {
        serverActions: {
          // edit: updated to new key. Was previously `allowedForwardedHosts`
          allowedOrigins: ['localhost'],
        },
    },
    basePath: `${process.env.BASE_PATH}`
}

module.exports = nextConfig
