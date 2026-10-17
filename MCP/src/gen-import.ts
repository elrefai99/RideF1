// @ts-nocheck — auto-generated barrel with lazy CJS re-exports
/**
 * gen-import.ts — AUTO-GENERATED, do not edit manually.
 * Regenerate: npx gen-import
 *
 * Value exports use lazy getters to prevent circular-dependency
 * errors when source files import from this barrel (CJS).
 *
 * Getters are installed on module.exports, not exports: esbuild-based
 * loaders (tsx, bun) reassign module.exports for any file containing
 * export syntax, which would strand getters bound to exports.
 */

export declare const dotenv: typeof import('./config/dotenv').default;
export declare const allUser: typeof import('./functions/user').allUser;
export declare const firstUser: typeof import('./functions/user').firstUser;
export declare const countAllUsers: typeof import('./functions/user').countAllUsers;
export declare const GET_LAST_USERS: typeof import('./tools/user').GET_LAST_USERS;
export declare const GET_FIRST_USERS: typeof import('./tools/user').GET_FIRST_USERS;
export declare const GET_COUNT_USERS: typeof import('./tools/user').GET_COUNT_USERS;

Object.defineProperty(module.exports, 'dotenv', { get() { return require('./config/dotenv').default }, enumerable: true, configurable: true });
Object.defineProperty(module.exports, 'allUser', { get() { return require('./functions/user').allUser }, enumerable: true, configurable: true });
Object.defineProperty(module.exports, 'firstUser', { get() { return require('./functions/user').firstUser }, enumerable: true, configurable: true });
Object.defineProperty(module.exports, 'countAllUsers', { get() { return require('./functions/user').countAllUsers }, enumerable: true, configurable: true });
Object.defineProperty(module.exports, 'GET_LAST_USERS', { get() { return require('./tools/user').GET_LAST_USERS }, enumerable: true, configurable: true });
Object.defineProperty(module.exports, 'GET_FIRST_USERS', { get() { return require('./tools/user').GET_FIRST_USERS }, enumerable: true, configurable: true });
Object.defineProperty(module.exports, 'GET_COUNT_USERS', { get() { return require('./tools/user').GET_COUNT_USERS }, enumerable: true, configurable: true });
