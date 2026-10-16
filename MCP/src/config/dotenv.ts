import dotenv from "dotenv";
import fs from 'node:fs'
import path from 'node:path'

const envPath = process.env.NODE_ENV === 'development' ? '.env.dev' : '.env';
const envFiles = envPath === '.env.dev' ? ['.env.dev', '.env'] : ['.env'];
const candidates = [
     ...envFiles.map((file) => path.resolve(process.cwd(), file)),
     ...envFiles.map((file) => path.resolve(process.cwd(), 'MCP', file)),
];
const resolvedEnvPath = candidates.find((candidate) => fs.existsSync(candidate));

export default dotenv.config({
     path: resolvedEnvPath ?? candidates[0]
})
