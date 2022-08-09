import Link from 'next/link';

export default function Login() {
	return (
		<div className="flex flex-col p-6 rounded-md sm:p-10 dark:bg-gray-900 dark:text-gray-100">
			<div className="mb-8 text-center">
				<h1 className="my-3 text-4xl font-bold">Sign in</h1>
				<p className="text-sm dark:text-gray-400">Sign in to access your account</p>
			</div>

			<form action="" className="space-y-12 ng-untouched ng-pristine ng-valid">
				<div className="space-y-4">
					<div>
						<label className="block mb-2 text-sm">Email address</label>
						<input type="email" name="email" id="email" placeholder="example@domain.com" className="w-full px-3 py-2 border rounded-md dark:border-gray-700 dark:bg-gray-900 dark:text-gray-100" />
					</div>

					<div>
						<div className="flex justify-between mb-2">
							<label className="text-sm">Password</label>
							<Link href="#">
								<a className="text-xs hover:underline dark:text-gray-400">Forgot password?</a>
							</Link>
						</div>

						<input type="password" name="password" id="password" placeholder="*****" className="w-full px-3 py-2 border rounded-md dark:border-gray-700 dark:bg-gray-900 dark:text-gray-100" />
					</div>
				</div>

				<div className="space-y-2">
					<div>
						<button type="button" className="w-full px-8 py-3 font-semibold rounded-md dark:bg-violet-400 dark:text-gray-900">Sign in</button>
					</div>
					<p className="px-6 text-sm text-center dark:text-gray-400">Don't have an account yet?
						<Link href="#">
							<a className="hover:underline dark:text-violet-400">Sign up</a>.
						</Link>
					</p>
				</div>
			</form>
		</div>
	);
}

