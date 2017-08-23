import React from "react";
import PropTypes from "prop-types";
import {withStyles} from "material-ui/styles";
import Select from "react-select";

const styles = {
    dropdown: {
        width: '100px',
        display: 'inline-block',
        verticalAlign: 'middle',
        fontSize: '18px'
    }
};

const buildTypes = [
    {
        id: "DOCKER",
        displayName: "Docker"
    },
    {
        id: "WAR",
        displayName: "WAR"
    }
    ];

class BuildTypeDropdown extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedBuildType: null
        };
    }

    selectedBuildType = (selectedBuildType) => {
        this.props.onBuildTypeSelection(selectedBuildType.value);
        this.setState({
            selectedBuildType: selectedBuildType
        });
    };

    mapBuildTypes = () => {
        return buildTypes.map(buildType => ({
            value: buildType.id,
            label: buildType.displayName
        }));
    };

    render() {
        const classes = this.props.classes;
        let buildTypesObjects = this.mapBuildTypes();
        return (
            <Select
                name="form-field-name"
                value={!this.state.selectedBuildType ? buildTypesObjects[0] : this.state.selectedBuildType}
                options={buildTypesObjects}
                onChange={this.selectedBuildType}
                clearable={false}
                className={classes.dropdown}
                searchable={false}
            />
        )
    }
}

BuildTypeDropdown.propTypes = {
    classes: PropTypes.object.isRequired,
    onBuildTypeSelection: PropTypes.func.isRequired
};

export default withStyles(styles)(BuildTypeDropdown);
