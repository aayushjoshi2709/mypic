import { Link } from 'react-router'
import type { GroupInterface } from '../../common/interfaces'

const GroupCard = ({groupData}: {groupData: GroupInterface}) => {
  return (
    <Link to={`/dashboard/groups/${groupData.id}`}>
      <div className="text-center w-fit">
          <div className="shadow w-[200px] h-[200px] rounded-full flex justify-center align-center items-center">
              <img className='w-full h-full rounded-full' src={groupData.imageUrl}/>
          </div>
          <p className='text-xl mt-2'>{groupData.name}</p>
      </div>
    </Link>
  )
}

export default GroupCard