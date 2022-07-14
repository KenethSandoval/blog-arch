import Link from "next/link";

const Header: React.FC = () => {
	return (
	  <header className="py-2">
		  <Link href="/">
				<a className="text-2xl font-bold text-green-500">Gestor de documentacion tecnica</a>
			</Link>
		</header>
	);
}

export default Header;
