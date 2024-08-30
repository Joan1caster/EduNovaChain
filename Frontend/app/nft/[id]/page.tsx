'use client';
import Alert, { ALERT_TYPE } from '@/app/components/Alert';
import { TRANSACTION_DATA, TRANSACTION_DATA_TYPE } from '@/app/constans';
import { VOTE_TYPE } from '@/app/types';
import { Button } from '@vercel/examples-ui';
import { useState } from 'react';
import Image from 'next/image';

export default function Page({ params: { id } }: { params: { id: string } }) {
  const data: TRANSACTION_DATA_TYPE = TRANSACTION_DATA.filter(
    (item) => item.cid === id
  )[0];
  const [message, setMessage] = useState('');
  const [type, setType] = useState<ALERT_TYPE>('info');

  const onBuyNow = () => {};

  const onLike = () => setMessage('Thank you for liking!');
  const onVote = (voteType: VOTE_TYPE) => {
    setMessage('Thank you for voting!');
  };

  const onReport = () => {
    setType('error');
    setMessage('Thank you for reporting!');
  };

  const onClose = () => {
    setType('info');
    setMessage('');
  };

  return (
    <div>
      <div className="grid grid-cols-2 md:grid-cols-1 gap-4 content-start">
        <Alert type={type} message={message} onClose={onClose} />
        <div className="aspect-h-1 aspect-w-1 w-full h-52 overflow-hidden flex justify-center items-center bg-gray-200 dark:bg-gray-600 lg:aspect-none group-hover:opacity-75 lg:h-60">
          <Image src={data.src} alt={data.name} width={50} height={52} />
        </div>
        <div>
          <h3 className="text-base/7 font-medium text-black ">{data.name}</h3>
          <p className="mt-2 text-sm/6 text-black ">{data.description}</p>
          <div className="mt-4 flex gap-2">
            <Button
              className="inline-flex items-center gap-2 rounded-md bg-indigo-600 py-1.5 px-3 text-sm/6 font-semibold text-white shadow-inner shadow-white/10 focus:outline-none data-[hover]:bg-indigo-400 data-[focus]:outline-1 data-[focus]:outline-white data-[open]:bg-indigo-500"
              onClick={onBuyNow}
            >
              Buy now
            </Button>
            <Button
              className="inline-flex items-center gap-2 rounded-md bg-red-600 py-1.5 px-3 text-sm/6 font-semibold text-white shadow-inner shadow-white/10 focus:outline-none data-[hover]:bg-red-400 data-[focus]:outline-1 data-[focus]:outline-white data-[open]:bg-red-500"
              onClick={onReport}
            >
              Report
            </Button>
          </div>
          <div className="mt-4 flex gap-2">
            <Button
              className="inline-flex items-center gap-2 rounded-md py-1.5 px-3 dark:text-white text-sm/6 font-semibold shadow-inner shadow-white/10 focus:outline-none data-[focus]:outline-1 data-[focus]:outline-white data-[open]:bg-indigo-500"
              onClick={onLike}
            >
              ❤️ {data.like}
            </Button>
            <Button
              className="inline-flex items-center gap-2 rounded-md py-1.5 px-3 dark:text-white text-sm/6 font-semibold shadow-inner shadow-white/10 focus:outline-none data-[focus]:outline-1 data-[focus]:outline-white data-[open]:bg-indigo-500"
              onClick={() => onVote('Support')}
            >
              👍 {data.support}
            </Button>
            <Button
              className="inline-flex items-center gap-2 rounded-md py-1.5 px-3 dark:text-white text-sm/6 font-semibold shadow-inner shadow-white/10 focus:outline-none data-[focus]:outline-1 data-[focus]:outline-white data-[open]:bg-indigo-500"
              onClick={() => onVote('Support')}
            >
              👎 {data.oppose}
            </Button>
            <Button
              className="inline-flex items-center gap-2 rounded-md py-1.5 px-3 dark:text-white text-sm/6 font-semibold shadow-inner shadow-white/10 focus:outline-none data-[focus]:outline-1 data-[focus]:outline-white data-[open]:bg-indigo-500"
              onClick={() => onVote('Support')}
            >
              👀 {data.abstain}
            </Button>
          </div>
        </div>
      </div>

      <div className="mt-2"></div>
    </div>
  );
}
