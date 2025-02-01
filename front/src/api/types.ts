
export type RequestTopicRecords  = {
    topics: string[]
    page: number
}

export type ResponseTopicRecords = {
    count_page: number
    records: string[]
}