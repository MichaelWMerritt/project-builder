import React from "react";
import PropTypes from "prop-types";
import {withStyles} from "material-ui/styles";
import Table, {TableBody, TableCell, TableHead, TableRow} from "material-ui/Table";
import {connect} from "react-redux";
import {bindActionCreators} from "redux";
import {loadReleases} from "../actions/release_actions";
import {loadModulesForRelease} from "../actions/module_actions";

export const RELEASES_VIEW_ID = 'RELEASES_VIEW';

const styles = theme => {

};

class ReleasesView extends React.Component {
    componentWillMount() {
        this.props.loadReleases();
    }

    renderReleases = () => {
        if (this.props.releases) {
            return this.props.releases.map((release) => {
                return (
                    <TableRow key={release._id}>
                        <TableCell>
                        {release.displayName}
                        </TableCell>
                        <TableCell>
                        {release.description}
                        </TableCell>
                        <TableCell>
                        {release.versionInfo.version}
                        </TableCell>
                        <TableCell>
                        {release.versionInfo.url}
                        </TableCell>
                        <TableCell>
                        {release.versionInfo.repoType}
                        </TableCell>
                    </TableRow>
                );
            });
        }
    };

    renderReleasesPanel = () => {
        if (this.props.releases && this.props.releases.length > 0) {
            return (
                <div style={{margin:'0', fontWeight:'normal'}}>
                    <Table>
                        <TableHead>
                            <TableRow>
                                <TableCell>Release Name</TableCell>
                                <TableCell>Description</TableCell>
                                <TableCell>Release</TableCell>
                                <TableCell>URL</TableCell>
                                <TableCell>Repo Type</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {this.renderReleases()}
                        </TableBody>
                    </Table>

                </div>
            )
        } else {
            return (
                <div style={{marginLeft:'20%',
                    marginRight:'20%',
                    marginTop: '25%',
                    marginBottom:'50%',
                    fontFamily:'"Varela Round",sans-serif',
                    fontSize:'20px',
                    fontWeight:'bold',
                    color:'#333'}}>
                    Define Releases in DB
                </div>
            )
        }
    };

    render() {
        return (
            <div>
                {this.renderReleasesPanel()}
            </div>
        )
    }
}

ReleasesView.propTypes = {
    classes: PropTypes.object.isRequired,
};

function mapPropsToState(state) {
    return {
        releases: state.releases,
        modules: state.modules
    }
};

function matchDispatchToProps(dispatch) {
    return bindActionCreators({
        loadReleases: loadReleases,
        loadModulesForRelease: loadModulesForRelease,
    }, dispatch);
}

export default withStyles(styles)(connect(mapPropsToState, matchDispatchToProps)(ReleasesView));
