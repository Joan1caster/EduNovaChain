'use client';
import Image from 'next/image';
import { Text, LoadingDots, Button } from '@vercel/examples-ui';
import { useState } from 'react';
import { UploadNft } from '@/app/components/UploadNft';
import { ConnectWallet } from '@/app/components/ConnectWallet';
import { SwitchNetwork } from '@/app/components/SwitchNetwork';

import Moralis from 'moralis-v1';

enum MintState {
  Connect,
  ConfirmNetwork,
  Upload,
  ConfirmMint,
  Loading,
}

export default function Page() {
  const [state, setState] = useState<MintState>(MintState.Connect);
  const [isLoading, setLoading] = useState(false);
  const [asset, setAsset] = useState<Moralis.File | null>(null);

  const onUploadComplete = async (asset: Moralis.File) => {
    setAsset(asset);
    setState(MintState.ConfirmMint);
    setLoading(false);
  };
  return (
    <div className="inline-block align-bottom text-left overflow-hidden  transform transition-all sm:my-8 sm:align-middle  ">
      {state === MintState.Connect && <ConnectWallet />}

      {state === MintState.ConfirmNetwork && <SwitchNetwork />}

      {state === MintState.Upload && <UploadNft onDone={onUploadComplete} />}

      {state === MintState.ConfirmMint && (
        <>
          <Text variant="h2">Confirm your mint</Text>
          <Text className="mt-6">
            Your image will be minted as an ERC721 Token. It can happen that
            images stored on IPFS as it is a distributed file hosting system
            that can fail. This is still the prefered method of choice to host
            in the NFT community as it is decentralized.{' '}
            <span className="underline italic">
              This process might take up to 1 minute to complete
            </span>
          </Text>
          {asset?._url ? (
            <section className="relative w-full pb-[20%] h-48 pb-6 mt-12">
              <Image
                className="rounded-xl"
                src={String(asset?._url)}
                alt="The image that will be minted as an NFT"
                layout="fill"
                objectFit="contain"
              />
            </section>
          ) : (
            <></>
          )}

          <section className="flex justify-center mt-6">
            <Button size="lg" variant="black" disabled={isLoading}>
              {isLoading ? <LoadingDots /> : 'Mint'}
            </Button>
          </section>
        </>
      )}
    </div>
  );
}
