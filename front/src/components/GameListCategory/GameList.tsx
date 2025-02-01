import './styles.scss'
import GameListItem from "./GameListItem.tsx";
import {useTopicList} from "@/api/getTopicList.ts";
import Loader from "@/components/Loader";

function GameList() {
    const {data, isLoading} = useTopicList();

    if (isLoading)
        return <Loader/>

    if (!data)
        return <div>Empty</div>

    return <ul className='game-list'>
        {
            data.map((topic) => (<GameListItem key={topic} topic={topic}>{topic}</GameListItem>))
        }
    </ul>
}

export default GameList;