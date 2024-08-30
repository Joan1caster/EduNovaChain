import type { Metadata } from 'next';
import type { AppProps } from 'next/app';
import { Layout, type LayoutProps } from '@vercel/examples-ui';
import Image from 'next/image';
import { Inter } from 'next/font/google';
import './globals.css';
import Link from 'next/link';
import LoginButton from './components/LoginButton';
import { cookies } from 'next/headers';
import WagmiContext from './context/WagmiContext';
import { ConnectButton } from '@rainbow-me/rainbowkit';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'Create Next App',
  description: 'Generated by create next app',
};

export default function RootLayout({
  children,
  modal,
}: Readonly<{
  children: React.ReactNode;
  modal: React.ReactNode;
}>) {
  const email = cookies().get('email')?.value ?? '';

  return (
    <html lang="en">
      <body className={inter.className}>
        <WagmiContext>
          <header className="px-2 py-4">
            <div className="flex justify-between">
              <Link href="/">
                {/* <Image
                src="/vercel.svg"
                alt="Vercel Logo"
                width={100}
                height={24}
                priority
              /> */}
                <img
                  className="mx-auto h-5 w-auto inline-block mr-2"
                  src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
                  alt="Your Company"
                />
                EduNovaChain
              </Link>

              <div className="flex">
                <LoginButton email={email} />
              </div>
            </div>
          </header>
          <hr className="border-t border-accents-2 border-t-gray-700 mb-6" />

          <section className='mx-auto max-w-2xl px-4 py-16 sm:px-6 sm:py-24 lg:max-w-7xl lg:px-8"'>
            <ConnectButton />
            {children}
            {modal}
          </section>

          <hr className="border-t border-accents-2 mt-6 border-t-gray-700" />

          <footer className="p-4 w-full">
            <div className="text-center text-sm">EduNovaChain © 2024</div>
          </footer>
        </WagmiContext>
      </body>
    </html>
  );
}
