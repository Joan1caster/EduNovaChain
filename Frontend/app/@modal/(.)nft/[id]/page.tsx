'use client';
import { Button, Dialog, DialogPanel, DialogTitle } from '@headlessui/react';
import { useState } from 'react';
import Image from 'next/image';
import { useRouter } from 'next/navigation';
import { TRANSACTION_DATA, TRANSACTION_DATA_TYPE } from '@/app/constans';
import { VOTE_TYPE } from '@/app/types';
import Alert, { ALERT_TYPE } from '@/app/components/Alert';

export default function Page({ params: { id } }: { params: { id: string } }) {
  const router = useRouter();
  const data: TRANSACTION_DATA_TYPE = TRANSACTION_DATA.filter(
    (item) => item.cid === id
  )[0];
  const [message, setMessage] = useState('');
  const [type, setType] = useState<ALERT_TYPE>('info');

  const onBuyNow = () => {};

  const onLike = () => setMessage('Thanks for your like!');
  const onVote = (voteType: VOTE_TYPE) => {
    setMessage('Thanks for your vote!');
  };

  const onReport = () => {
    setType('error');
    setMessage('Thanks for your report!');
  };

  const onClose = () => {
    setType('info');
    setMessage('');
  };
  return (
    <div className="fixed top-0 left-0">
      <Dialog
        as="div"
        open
        className="fixed inset-0 flex w-screen items-center justify-center bg-black/30 p-4 transition duration-300 ease-out data-[closed]:opacity-0"
        onClose={() => router.back()}
      >
        <Alert type={type} message={message} onClose={onClose} />
        <DialogPanel
          transition
          className="w-full max-w-md rounded-xl p-6 bg-white dark:bg-black border border-slate-200 duration-300 ease-out data-[closed]:transform-[scale(95%)] data-[closed]:opacity-0"
        >
          <DialogTitle
            as="h3"
            className="text-base/7 font-medium text-black dark:text-white"
          >
            {data.name}
          </DialogTitle>

          <div className="mt-2 aspect-h-1 aspect-w-1 w-full h-52 overflow-hidden flex justify-center items-center bg-gray-200 dark:bg-gray-600 lg:aspect-none group-hover:opacity-75 lg:h-60">
            <Image src={data.src} alt={data.name} width={50} height={52} />
          </div>
          <p className="mt-2 text-sm/6 text-black dark:text-white">
            {data.description}
          </p>
          <div className="mt-4 flex justify-between">
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
        </DialogPanel>
      </Dialog>
    </div>
  );
}
