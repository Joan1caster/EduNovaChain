import Banner from "./components/home/Banner";
import BestSellerList from "./components/home/BestSellerList";
import HotSellerList from "./components/home/HotSellerList";
import Suggestion from "./components/home/Suggestion";
import Topic from "./components/home/Topic";

export default function Home() {
  return (
    <>
      <Banner />
      <Suggestion />
      <BestSellerList />
      <HotSellerList />
      <Topic />
    </>
  );
}
