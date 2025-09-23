import NewsCard from "./components/NewsCard";

export type News = {
  slug: string;
  headline: string;
  publishedOn: string;
  imageUrl: string | null;
  newsUrl: string;
};

export default async function Home() {
  const response = await fetch("http://go.learn/news");
  const data: { status: string; data: News[] } | null = response.ok
    ? await response.json()
    : null;
  return (
    <main className="size-full">
      <div className="flex justify-center items-center h-64">
        <h1>News</h1>
      </div>
      <div className="columns-5 space-y-5 p-10">
        {data?.data &&
          data?.data.map((news) => <NewsCard key={news.slug} {...news} />)}
      </div>
    </main>
  );
}
