import React from "react";
import PropTypes from "prop-types";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import Toolbar from "material-ui/Toolbar";
import Button from "material-ui/Button";
import IconButton from "material-ui/IconButton";
import Icon from "material-ui/Icon";
import Typography from "material-ui/Typography";
import {withStyles} from "material-ui/styles";
import Badge from "material-ui/Badge";
import ReactTooltip from "react-tooltip";
import {selectView} from "../actions/view_selection";
import {BUILD_REFERENCES_VIEW_ID} from "../containers/build_references_view.jsx";
import {RELEASES_VIEW_ID} from "../containers/releases_view.jsx";
import {BUILD_CREATION_VIEW_ID} from "../containers/build_creation_view.jsx";

const styles = {
    root: {
        position: 'fixed',
        height: '100px',
        backgroundColor: '#2D3E50',
        width: '100%',
        borderTopColor: '#89C649',
        borderTopWidth: '3px',
        borderTopStyle: 'solid',
        zIndex: 100000
    },
    toolbar: {
        marginLeft:'20%',
        marginRight:'20%',
        paddingLeft:'0px',
        paddingRight:'0px',
        height:'97px',
        backgroundColor:'#2D3E50'
    },
    typography: {
        fontSize:'2.0vw',
        fontFamily:'"Varela Round",sans-serif',
        textTransform:'uppercase',
        verticalAlign:'middle',
        color:'#F9F9F9',
        flex: 1
    },
    iconButton: {
        width: '30px',
        height: '40px',
        marginTop: '7px',
    },
    iconHover: {
        color: '#F9F9F9',
        '&:hover': {
            color: '#89C649',
        },
    },
    button: {
        fontSize: '20px',
        fontFamily:'"Varela Round",sans-serif',
        marginRight: '15px',
        marginTop: '5px',
        textTransform: 'capitalize',
        color: '#F9F9F9',
        '&:hover': {
            color: '#89C649',
        },
    },
    badge: {
        marginTop: '10px',
        marginRight: '5px'
    }
};

class Header extends React.Component {
    selectView = (viewId) => {
        this.props.selectView(viewId);
    };

    render() {
        const classes = this.props.classes;
        return (
            <div className={classes.root}>
                <Toolbar className={classes.toolbar}>
                    <Typography className={classes.typography} type="title" color="inherit">
                        Project Builder
                    </Typography>
                    <Button className={classes.button} onClick={() => this.selectView(BUILD_CREATION_VIEW_ID)}>
                        Create New Build
                        <Icon style={{fontSize: '24px', marginLeft:'5px'}}>add</Icon>
                    </Button>
                    <hr style={{color: '#757575', height: '38px', borderStyle:'solid', backgroundColor: '#757575', marginRight: '20px'}} />
                    <IconButton data-tip="Releases" style={{marginTop: '3px', marginRight: '20px'}} className={classes.iconHover} onClick={() => this.selectView(RELEASES_VIEW_ID)} >
                        <Icon style={{fontSize: '30px'}}>history</Icon>
                    </IconButton>
                    <IconButton data-tip="Builds" style={{marginTop: '8px'}} className={classes.iconHover} onClick={() => this.selectView(BUILD_REFERENCES_VIEW_ID)} >
                        <Badge badgeContent={this.props.buildReferences.length} color="primary">
                            <Icon style={{fontSize:'30px'}} className={classes.iconHover}>build</Icon>
                        </Badge>
                    </IconButton>
                </Toolbar>
                <ReactTooltip place="bottom" />
            </div>
        )
    }
}

Header.propTypes = {
    classes: PropTypes.object.isRequired,
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

export default withStyles(styles)(connect(mapStateToProps, matchDispatchToProps)(Header));

