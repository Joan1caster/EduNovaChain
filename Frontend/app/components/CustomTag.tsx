export const AiTag = () => {
  return (
    <div className="w-6 h-4 inline-block bg-[url('/images/slice/ai_tag.png')] bg-contain bg-no-repeat text-center text-sm text-white ml-1"></div>
  );
};

export const OrderTag = ({
  order,
  bg = true,
}: {
  order: number;
  bg?: boolean;
}) => {
  return (
    <div
      className={`w-6 h-6 inline-block bg-[url('/images/slice/${order > 3 && bg ? "n" : order}.png')] bg-contain bg-no-repeat text-center text-sm ${bg ? "text-[#333]" : "text-[#999]"} leading-6`}
    >
      {order}
    </div>
  );
};

export const BannerCard = ({
  order,
  children,
}: {
  order: number;
  children: React.ReactNode;
}) => {
  return (
    <div
      className={`flex-none w-[500px] h-[232px] pr-3 bg-[url('/images/slice/banner_bg_${order % 3 + 1}.png')] bg-contain bg-no-repeat`}
    >
      {children}
    </div>
  );
};

export const TopicCard = ({
  order,
  children,
}: {
  order: number;
  children: React.ReactNode;
}) => {
  return (
    <div
      className={`relative w-full h-full bg-[url('/images/slice/card_${order + 1}.png')] bg-contain bg-no-repeat`}
    >
      {children}
    </div>
  );
};
