import Image from "next/image";
import { News } from "../page";

export default function NewsCard({
  headline,
  imageUrl = "https://www.un.org/sites/un2.un.org/files/styles/large-article-image-style-16-9/public/field/image/775566_1.jpg",
  publishedOn,
}: Omit<News, "slug">) {
  return (
    <div className="grid grid-cols-[auto_1fr] bg-slate-100 text-foreground rounded-2xl p-2 overflow-hidden gap-2">
      <div className="relative w-full aspect-square row-span-2">
        <Image
          fill
          alt={headline}
          src={imageUrl ?? "/placeholder.png"}
          className="rounded-xl"
          objectFit="cover"
        />
      </div>
      <div className={"line-clamp-2 text-xs"}>{headline}</div>
      <span className="text-xs text-zinc-500">{publishedOn}</span>
      {/*<div className="p-0 flex justify-between">*/}
      {/*  <a*/}
      {/*    href={newsUrl}*/}
      {/*    className="bg-primary rounded-full p-2 px-4 text-sm text-background"*/}
      {/*  >*/}
      {/*    Read Full*/}
      {/*  </a>*/}
      {/*</div>*/}
    </div>
  );
}
