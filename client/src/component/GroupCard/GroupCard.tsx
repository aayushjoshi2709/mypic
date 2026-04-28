import type { GroupListInterface } from '../../common/interfaces'

interface GroupCardInterface{
  groupData: GroupListInterface,
  onClick: (id:string) => void
}

const GroupCard = ({groupData, onClick}: GroupCardInterface) => {
  return (
      <div onClick={()=>onClick(groupData.id)} className="text-center w-fit">
          <div className="shadow w-[200px] h-[200px] rounded-full flex justify-center align-center items-center">
              <img className='w-full h-full rounded-full' src={groupData.imageUrl}/>
          </div>
          <p className='text-xl mt-2'>{groupData.name}</p>
      </div>
  )
}

export default GroupCard