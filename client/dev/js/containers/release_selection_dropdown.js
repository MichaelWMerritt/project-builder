import React from "react";
import PropTypes from "prop-types";
import {withStyles} from "material-ui/styles";
import {connect} from "react-redux";
import Select from "react-select";

const styles = {
    dropdown: {
        width: '100px',
        display: 'inline-block',
        verticalAlign: 'middle',
        fontSize: '18px'
    }
};

class ReleaseSelectionDropdown extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedRelease: null
        };
        this.mapReleases = this.mapReleases.bind(this);
        this.selectedRelease = this.selectedRelease.bind(this);
    }

    selectedRelease(selectedRelease) {
        this.props.onReleaseSelection(selectedRelease.value);
        this.setState({
            selectedRelease: selectedRelease
        });
    }

    mapReleases() {
        return this.props.releases.map(release => ({
            value: release._id,
            label: release.displayName
        }));
    }

    render() {
        const classes = this.props.classes;
        if (!this.props.releases) {
            return null
        }
        let releases = this.mapReleases();
        return (
            <Select
                name="form-field-name"
                value={!this.state.selectedRelease ? releases[0] : this.state.selectedRelease}
                options={releases}
                onChange={this.selectedRelease}
                clearable={false}
                className={classes.dropdown}
                searchable={false}
            />
        )
    }
}

ReleaseSelectionDropdown.propTypes = {
    classes: PropTypes.object.isRequired,
    onReleaseSelection: PropTypes.func.isRequired
};

function mapStateToProps(state) {
    return {
        releases: state.releases
    };
}

export default withStyles(styles)(connect(mapStateToProps)(ReleaseSelectionDropdown));