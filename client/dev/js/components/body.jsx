import React, {PropTypes} from "react";
import {withStyles} from "material-ui/styles";
import {connect} from "react-redux";
import BuildCreationForm, {BUILD_CREATION_VIEW_ID} from "../containers/build_creation_view.jsx";
import BuildReferencesView, {BUILD_REFERENCES_VIEW_ID} from "../containers/build_references_view.jsx";
import ReleasesView, {RELEASES_VIEW_ID} from "../containers/releases_view.jsx";

const styles = {
    body: {
        height:'-webkit-fill-available',
        paddingTop: '130px',
        marginLeft: '20%',
        marginRight: '20%',
    }
};

class Body extends React.Component {
    renderView = () => {
        switch (this.props.activeView) {
            case BUILD_CREATION_VIEW_ID:
                return (
                    <BuildCreationForm ref="buildCreationForm"/>
                );
            case RELEASES_VIEW_ID:
                return (
                    <ReleasesView />
                );
            case BUILD_REFERENCES_VIEW_ID:
            default:
                return (
                    <BuildReferencesView />
                );
        }
    };

    render() {
        const classes = this.props.classes;
        return (
            <div className={classes.body}>
                {this.renderView()}
            </div>
        );
    }
}

Body.propTypes = {
    classes: PropTypes.object.isRequired,
};

function mapStateToProps(state) {
    return {
        activeView: state.activeView
    };
}

export default withStyles(styles)(connect(mapStateToProps)(Body));