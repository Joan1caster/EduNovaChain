import { Text, Page } from '@vercel/examples-ui';
import Link from 'next/link';
import TransactionList from './components/TransactionList';
// import { Mint } from '../components/Mint'

export default function Home() {
  return (
    <div>
      <section className="flex flex-col gap-6">
        <div className="flex justify-between">
          <div>
            <Text>NFTs</Text>
            <Text>
              A list of all the NFTs including their image, name, description,
              CID.
            </Text>
          </div>

          <Link href="/nft/create">
            <div className="rounded-md px-4 py-2 text-xs text-white bg-indigo-600 shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
              Create NFT
            </div>
          </Link>
        </div>

        <TransactionList />
      </section>

      <section className="flex flex-col ">{/* <Mint /> */}</section>
    </div>
  );
}
