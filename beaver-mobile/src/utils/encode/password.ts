import { MD5 } from 'crypto-js';

export function encodePassword(password: string, salt: string): string {
	const passwordSalted = salt + password;
	return MD5(passwordSalted).toString()
}
