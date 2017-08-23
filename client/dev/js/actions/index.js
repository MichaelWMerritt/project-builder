//action creator
export const selectUser = (user) => {
    //action
    return {
        //type
        type: "USER_SELECTED",
        //payload
        payload: user
    }
};