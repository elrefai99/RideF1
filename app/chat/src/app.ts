import './config/dotenv.conf'
import fasti from 'fastify'
import cors from '@fastify/cors'

const app = fasti({
     logger: true
})

async function start() {
     await app.register(cors, {
          origin: process.env.SITE_DOMAIN as string,
          credentials: true
     })

     app.get('/health', () => (
          {
               status: "OK"
          }
     ))

     await app.listen({
          port: 9001,
          host: "0.0.0.0",
     }).then(() => {
          console.log("Chat server runner")
     }).catch(e => console.error(e));
}

start()
