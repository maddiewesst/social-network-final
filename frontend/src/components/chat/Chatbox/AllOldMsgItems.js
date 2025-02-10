import React from "react";
import OldMsgItem from "./OldMsgItem";

const AllOldMsgItems = (props) => {
    
    return (
        props.oldMsgItems.map((oldMsg) => {
            return <OldMsgItem
                key={oldMsg.id}
                id={oldMsg.id}
                targetid={oldMsg.targetid}
                sourceid={oldMsg.sourceid}
                // groupid={oldMsg.groupid}
                msg={oldMsg.message}
                createdat={oldMsg.createdat}
            />
        })
    );
}

export default React.memo(AllOldMsgItems);