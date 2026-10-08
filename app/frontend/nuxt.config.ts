export default defineNuxtConfig({
  nitro: {
    externals: {
      inline: [/[\\/]node_modules[\\/]nuxt[\\/]dist[\\/]/],
    },
  },
  runtimeConfig: {
    public: {
      apiBase: process.env.NUXT_PUBLIC_API_BASE || 'http://localhost:8080',
    },
  },
})
