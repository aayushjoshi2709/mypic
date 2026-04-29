import { useDispatch, useSelector } from "react-redux";
import { setModal } from "../../store/modal.slice";
import { ModalNames } from "../../common/Constants";
import AddGroupModal from "../../component/Modal/AddGroupModal/AddGroupModal";
import type { RootState } from "../../store/store";
import { useEffect } from "react";
import { setFetchGroups } from "../../store/group.slice";
import GroupCard from "../../component/GroupCard/GroupCard";
import { RoundedButtonSecondary } from "../../component/Button/RoundedButton";
import { useNavigate } from "react-router";

const Groups = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const group = useSelector((state: RootState) => state.group);

  useEffect(() => {
    if (group.groups == null) {
      dispatch(setFetchGroups());
    }
  }, [dispatch, group.groups]);

  return (
    <div>
      <AddGroupModal />
      {
        <div className="w-full">
          <div className="p-8 m-4 d-flex rounded-2xl content-center text-center border-2 border-dashed border-gray-400 bg-blue-100">
            <h1 className="text-4xl font-bold mb-4">Groups</h1>
            <h2 className="text-2xl mb-4">
              {(group.groups?.length ?? 0) === 0
                ? "Hey, Now you can group your pictures..."
                : "A lot of friend groups was waiting for you to visit."}
            </h2>
            {(group.groups?.length ?? 0) > 0 && (
              <h3 className="mb-2">View them below or</h3>
            )}
            <RoundedButtonSecondary
              text="Create new group"
              onClick={() => {
                dispatch(
                  setModal({
                    name: ModalNames.ADD_GROUP,
                  }),
                );
              }}
            />
          </div>

          <main className="p-8 m-4 flex flex-wrap gap-5 max-w-[calc(100%-1rem)]">
            {group.groups?.map((group) => {
              return (
                <GroupCard
                  key={group.id}
                  groupData={group}
                  onClick={(groupId) => {
                    navigate(`/dashboard/groups/${groupId}`);
                  }}
                />
              );
            })}
          </main>
        </div>
      }
    </div>
  );
};

export default Groups;
