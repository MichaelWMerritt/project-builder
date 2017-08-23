import React from "react";
import "react-select/dist/react-select.css";
import {createMuiTheme, MuiThemeProvider, withStyles} from "material-ui/styles";
import createPalette from "material-ui/styles/palette";
import {bindActionCreators} from "redux";
import {connect} from "react-redux";
import PropTypes from "prop-types";
import Header from "./header";
import Body from "./body.jsx";
import {loadAllBuildReferences} from "../actions/builds_actions";
require('../../scss/style.scss');

const theme = createMuiTheme({
    palette: createPalette({
        type: 'light'
    })
});

const styles = {
    app: {
        backgroundColor: '#FFF'
    }
};

class App extends React.Component {
    componentWillMount() {
        this.props.loadAllBuildReferences();
    }

    render() {
        const classes = this.props.classes;
        return (
            <MuiThemeProvider theme={theme}>
                <div className={classes.app}>
                    <Header />
                    <Body />
                </div>
            </MuiThemeProvider>
        )
    }
}

App.propTypes = {
    classes: PropTypes.object.isRequired,
};

function mapStateToProps(state) {
    return {
    };
}

function matchDispatchToProps(dispatch) {
    return bindActionCreators({
        loadAllBuildReferences: loadAllBuildReferences
    }, dispatch);
}

export default withStyles(styles)(connect(mapStateToProps, matchDispatchToProps)(App));