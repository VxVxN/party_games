import {Outlet} from "react-router";
import './styles.scss'

function MainLayout() {

    return <div className='main-layout'>
        <Outlet/>
    </div>
}

export default MainLayout;