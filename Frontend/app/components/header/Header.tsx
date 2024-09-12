import Image from "next/image";
import Link from "next/link";
import LoginButton from "../LoginButton";
import Search from "./Search";
import Nav from "./Nav";

export default function Header() {
  return (
    <header className="px-10 py-4 mb-1">
      <div className="flex justify-between items-center text-primary">
        <Link href="/">
          <Image
            src="/images/slice/logo_01.png"
            alt="Vercel Logo"
            width={250}
            height={64}
            priority
          />
        </Link>
        <Nav />
        <Search />
        <div className="flex">
          <LoginButton email={""} />
        </div>
      </div>
    </header>
  );
}
