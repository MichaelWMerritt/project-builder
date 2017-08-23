import React from "react";
import PropTypes from "prop-types";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import Toolbar from "material-ui/Toolbar";
import Button from "material-ui/Button";
import {withStyles} from "material-ui/styles";
import {selectView} from "../actions/view_selection";
import {BUILD_REFERENCES_VIEW_ID} from "../containers/build_references_view.jsx";

const styles = {
    cancelButton: {
        backgroundColor: '#EF5350',
        color: '#F9F9F9',
        fontSize: '20px'
    },
    creationButton: {
        backgroundColor: '#89C649',
        '&:disabled': {
            backgroundColor: '#C8E6C9',
        },
        float:'right',
        color: '#F9F9F9',
        fontSize: '20px'
    },
    root: {
        position: 'fixed',
        left:'0',
        bottom:'0',
        right:'0',
        height: '80px',
        backgroundColor: '#2D3E50',
        width: '100%',
        borderBottomColor: '#89C649',
        borderBottomWidth: '3px',
        borderBottomStyle: 'solid',
        zIndex: 100000
    },
    toolbar: {
        marginLeft:'20%',
        marginRight:'20%',
        paddingLeft:'0px',
        paddingRight:'0px',
        height:'80px',
        backgroundColor:'#2D3E50'
    },
    toolbarDiv: {
        display:'inline-block',
        width:'100%'
    }
};

class BuildCreationFormFooter extends React.Component {
    selectView = (viewId) => {
        console.log('view: ' + viewId);
        this.props.selectView(viewId);
        this.props.onCancel();
    };

    render() {
        const classes = this.props.classes;
        return (
            <div className={classes.root}>
                <Toolbar className={classes.toolbar}>
                    <div className={classes.toolbarDiv}>
                        <Button raised className={classes.cancelButton} onClick={() => this.selectView(BUILD_REFERENCES_VIEW_ID)}>
                            Cancel
                        </Button>
                        <Button raised className={classes.creationButton} disabled={!this.props.buildFormComplete} onClick={this.props.onCreate}>
                            Create
                        </Button>
                    </div>
                </Toolbar>
            </div>
        )
    }
}

BuildCreationFormFooter.propTypes = {
    classes: PropTypes.object.isRequired,
    onCancel: PropTypes.func.isRequired,
    onCreate: PropTypes.func.isRequired,
    buildFormComplete: PropTypes.bool.isRequired
};

function mapStateToProps(state) {
    return {
        activeView: state.activeView,
        buildReferences: state.buildReferences
    }
}

function matchDispatchToProps(dispatch) {
    return bindActionCreators({selectView: selectView}, dispatch);
}

export default withStyles(styles)(connect(mapStateToProps, matchDispatchToProps)(BuildCreationFormFooter));
