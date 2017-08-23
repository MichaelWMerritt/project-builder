import React from "react";
import PropTypes from "prop-types";
import {withStyles} from "material-ui/styles";
import Table, {TableBody, TableCell, TableHead, TableRow} from "material-ui/Table";
import {connect} from "react-redux";

export const BUILD_REFERENCES_VIEW_ID = 'BUILD_REFERENCES_VIEW';

const styles = theme => {

};

class BuildReferencesView extends React.Component {
    buildReferences = this.props.buildReferences.map((buildReference) => {
        return (
            <TableRow key={buildReference._id}>
                <TableCell>
                    {buildReference.buildReference.displayName}
                </TableCell>
                <TableCell>
                    {buildReference.status}
                </TableCell>
                <TableCell>
                    {buildReference.buildReference.release._id}
                </TableCell>
                <TableCell>
                    {buildReference.modules}
                </TableCell>
                <TableCell>
                    {buildReference.buildReference.buildType}
                </TableCell>
                <TableCell>
                    {buildReference.buildReference.dateCreated}
                </TableCell>
            </TableRow>
        );
    });

    renderBuildReferencesPanel = () => {
        if (this.props.buildReferences && this.props.buildReferences.length > 0) {
            return (
                <div>
                    <Table>
                        <TableHead>
                            <TableRow>
                                <TableCell>Build Name</TableCell>
                                <TableCell>Status</TableCell>
                                <TableCell>Release</TableCell>
                                <TableCell>Modules</TableCell>
                                <TableCell>Build Type</TableCell>
                                <TableCell>Date Created</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {this.buildReferences}
                        </TableBody>
                    </Table>
                </div>
            )
        } else {
            return (
                <h1 style={{textAlign:'center'}}>
                    No Builds
                </h1>
            )
        }
    };

    render() {
        return (
            <div style={{marginLeft:'20%',
                marginRight:'20%',
                marginTop: '25%',
                marginBottom:'50%',
                fontFamily:'"Varela Round",sans-serif',
                fontSize:'20px',
                fontWeight:'bold',
                color:'#333'}}>
                {this.renderBuildReferencesPanel()}
            </div>
        )
    }
}

BuildReferencesView.propTypes = {
    classes: PropTypes.object.isRequired,
};

function mapPropsToState(state) {
    return {
        buildReferences: state.buildReferences
    }
};

export default withStyles(styles)(connect(mapPropsToState)(BuildReferencesView));
