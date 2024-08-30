import Image from 'next/image';
import Link from 'next/link';

const data = [
  {
    src: '/next.svg',
    name: '123123',
    description: 'sheizhidaone',
    cid: '18971478236482',
    status: 'active',
  },
  {
    src: '/next.svg',
    name: '123123',
    description: 'sheizhidaone',
    cid: '18971478236483',
    status: 'active',
  },
  {
    src: '/next.svg',
    name: '123123',
    description: 'sheizhidaone',
    cid: '18971478236484',
    status: 'active',
  },
  {
    src: '/next.svg',
    name: '123123',
    description: 'sheizhidaone',
    cid: '18971478236485',
    status: 'active',
  },
  {
    src: '/next.svg',
    name: '123123',
    description: 'sheizhidaone',
    cid: '18971478236486',
    status: 'active',
  },
  {
    src: '/next.svg',
    name: '123123',
    description: 'sheizhidaone',
    cid: '18971478236487',
    status: 'active',
  },
  {
    src: '/next.svg',
    name: '123123',
    description: 'sheizhidaone',
    cid: '18971478236488',
    status: 'active',
  },
  {
    src: '/next.svg',
    name: '123123',
    description: 'sheizhidaone',
    cid: '18971478236489',
    status: 'active',
  },
];

export default function TransactionList() {
  return (
    <div className="mt-6 grid grid-cols-1 gap-4 sm:grid-cols-4 lg:grid-cols-5">
      {data.map((item) => (
        <div key={item.cid} className="bg-slate-50 rounded-md">
          <Link
            className="w-full h-full cursor-pointer"
            href={`/nft/${item.cid}`}
          >
            <div className="aspect-h-1 aspect-w-1 w-full h-40 overflow-hidden flex justify-center items-center rounded-tl-md rounded-tr-md bg-gray-200 lg:aspect-none group-hover:opacity-75 lg:h-60">
              <Image src={item.src} alt={item.name} width={50} height={50} />
            </div>
          </Link>
          <div className="mt-2 px-2">
            <h3 className="text-sm text-gray-700">{item.name}</h3>
          </div>
          <div className="mt-2 px-2">
            <p className="text-sm font-medium text-gray-900">$35</p>
          </div>
          <div className="mt-2 py-1 text-xs text-center text-white rounded-bl-md rounded-br-md cursor-pointer bg-indigo-600 shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">
            Buy now
          </div>
        </div>
      ))}
    </div>
  );
}
