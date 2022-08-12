import Link from 'next/link';
import { useState } from 'react';
import { login } from '@service/auth.service';
import { LoginData } from '@interface/auth.response';

export default function Login() {
	const [username, setUsername] = useState('');
	const [password, setPassword] = useState('');

	const handleSubmit = async (e: any) => {
		e.preventDefault();
		const payload: LoginData = {
				username,
				password
		}
	  await login(payload);	
	}

	return (
		<div className="flex flex-col p-6 rounded-md sm:p-10 dark:bg-gray-900 dark:text-gray-100">
			<div className="mb-8 text-center">
				<h1 className="my-3 text-4xl font-bold">Sign in</h1>
				<p className="text-sm dark:text-gray-400">Sign in to access your account</p>
			</div>

			<form className="space-y-12 ng-untouched ng-pristine ng-valid" onSubmit={handleSubmit}>
				<div className="space-y-4">
					<div>
						<label className="block mb-2 text-sm">Username</label>
						<input type="text" name="username" id="username" placeholder="Username" className="w-full px-3 py-2 border rounded-md dark:border-gray-700 dark:bg-gray-900 dark:text-gray-100" onChange={event => setUsername(event.target.value)}/>
					</div>

					<div>
						<div className="flex justify-between mb-2">
							<label className="text-sm">Password</label>
							<Link href="/password">
								<a className="text-xs hover:underline dark:text-gray-400">Forgot password?</a>
							</Link>
						</div>

						<input type="password" name="password" id="password" placeholder="*****" className="w-full px-3 py-2 border rounded-md dark:border-gray-700 dark:bg-gray-900 dark:text-gray-100" onChange={event => setPassword(event.target.value)}/>
					</div>
				</div>

				<div className="space-y-2">
					<div>
						<button type="submit" className="w-full px-8 py-3 font-semibold rounded-md dark:bg-violet-400 dark:text-gray-900">Sign in</button>
					</div>
					<p className="px-6 text-sm text-center dark:text-gray-400">Don't have an account yet?
							<a className="hover:underline dark:text-violet-400">Sign up</a>.
					</p>
				</div>
			</form>
		</div>
	);
}

