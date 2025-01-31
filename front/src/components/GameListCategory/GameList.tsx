import './styles.scss'
import GameListItem from "./GameListItem.tsx";
import {useTopicList} from "../../api/getTopicList.ts";
import Loader from "@/components/Loader";

function GameList() {
    const {data, isLoading} = useTopicList();

    if (isLoading)
        return <Loader/>

    return <ul className='game-list'>
        <GameListItem>+18</GameListItem>
        <GameListItem>middle</GameListItem>
        <GameListItem>hard</GameListItem>
    </ul>
}

export default GameList;