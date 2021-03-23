export type Tweet = {
    authorID: string,
    username: string,
    name: string,
    date: Date,
    text: string,
    profilePicURL: string,
    medias: Media[],
    refTweet: RefTweet[],
    likes: number,
    quotes: number,
    retweets: number,
}

export type RefTweet = {
    authorID: string,
    username: string,
    name: string,
    date: Date,
    text: string,
    profilePicURL: string,
}

export type Media = {
    width: number,
    previewImageURL: string,
    type: string,
    height: number,
    url: string,
}

export type List = {
    owner: string,
    searchterms: Search[],
    createdAt: Date,
    updatedAt: Date,
}

export type Search = {
    name: string,
    searchterm: string,
}

export type Error = {
    error: string;
}

export type TweetObject = {
    title: string,
    tweets: Tweet[],
}