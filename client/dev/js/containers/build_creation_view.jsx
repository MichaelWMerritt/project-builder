import React from "react";
import PropTypes from "prop-types";
import {withStyles} from "material-ui/styles";
import {bindActionCreators} from "redux";
import {connect} from "react-redux";
import {loadReleases} from "../actions/release_actions";
import {loadModulesForRelease} from "../actions/module_actions";
import ReleaseSelectionDropdown from "./release_selection_dropdown";
import {Column, Grid, Row} from "react-cellblock";
import Input from "material-ui/Input";
import InputLabel from "material-ui/Input/InputLabel";
import FormControl from "material-ui/Form/FormControl";
import IntegrationAutosuggest from "./module_suggestion_container.jsx";
import IconButton from "material-ui/IconButton";
import Icon from "material-ui/Icon";
import ReactTooltip from "react-tooltip";
import BuildCreationFormFooter from "../components/build_creation_form_footer";
import {createBuild, updateBuildFormState} from "../actions/builds_actions";
import {selectView} from "../actions/view_selection";
import {BUILD_REFERENCES_VIEW_ID} from "../containers/build_references_view.jsx";
import BuildTypeDropdown from "../components/build_type_dropdown.jsx";

export const BUILD_CREATION_VIEW_ID = 'BUILD_CREATION_VIEW';

const styles = {
    releaseSelectionLabel: {
        display: 'inline-block',
        verticalAlign: 'middle',
        fontFamily:'"Varela Round",sans-serif',
        margin: '0.3em',
        fontSize:'30px',
        fontWeight:'bold',
        color:'#333'
    },
    releaseSelection: {

    },
    h2: {
        fontFamily:'"Varela Round",sans-serif',
        marginTop:'10px',
        fontWeight: '400'
    },
    textField: {
        width: 300,
        color: '#333'
    },
    textFieldLabel: {
        color: '#333'
    },
    formControl: {
        marginTop:'30px',
        paddingLeft:'10px',
        width:'80%'
    },
    inputLabel: {
        paddingLeft:'10px'
    },
    iconHover: {
        color: '#757575',
        '&:hover': {
            color: '#333',
        },
    }
};

class BuildCreationForm extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedRelease: null,
            selectedModules: [],
            build: {
                buildType: null,
                utilityId: "",
                masterModelId: "",
                projectName: "",
                description: "",
                timezone: "America/New_York",
                defaultLatitude: "39.025585",
                defaultLongitude: "-77.073898",
                release: null,
                modules: []
            }
        };
    }

    componentWillMount() {
        this.props.loadReleases();
        this.updateBuildFormCompletion();
    }

    onReleaseSelection = (selectedRelease) => {
        this.setState((state) => {
            this.props.releases.map((release) => {
                if (release._id === selectedRelease) {
                    state.build["release"] = release
                }
            });
            state.selectedRelease = selectedRelease;
        });
        this.props.loadModulesForRelease(selectedRelease);
    };

    onModuleSelection = (selectedModule) => {
        let selectedModules = this.state.selectedModules;
        if (this.state.selectedModules.indexOf(selectedModule) > -1) {
            selectedModules.splice(this.state.selectedModules.indexOf(selectedModule), 1)
        } else {
            selectedModules.push(selectedModule)
        }
        this.setState((state) => {
            state.selectedModules = selectedModules;
            state.build["modules"] = selectedModules;
        });
        this.updateBuildFormCompletion();
    };

    onBuildTypeSelection = (buildType) => {
        this.setState((state) => state.build["buildType"] = buildType);
        this.updateBuildFormCompletion();
    };

    removeModuleFromSelectedModules = (moduleId) => {
        this.setState({
            selectedModules: this.state.selectedModules.filter((module) => module._id !== moduleId)
        });
        this.updateBuildFormCompletion();
    };

    createBuild = () => {
        this.validateBuildParameters();
        this.props.createBuild(this.state.build);
        this.props.selectView(BUILD_REFERENCES_VIEW_ID);
    };

    cancelBuild = () => {
        //Do cleanup here?
    };

    validateBuildParameters = () => {

    };

    updateBuildFormCompletion = () => {
        let buildFormComplete = this.state.selectedModules.length > 0 && this.state.build.utilityId !== "";
        this.props.updateBuildFormState(buildFormComplete);
    };

    validateRequiredField = (event) => {
        event.persist();
        this.setState((state) => {
            if (state.build["release"] === null) {
                state.build["release"] = this.props.releases[0];
            }
            state.build[event.target.name] = event.target.value;
        });
        this.updateBuildFormCompletion();
    };

    render() {
        const classes = this.props.classes;
        let selectedModules = null;
        if (this.state.selectedModules.length > 0) {
            selectedModules = this.state.selectedModules.map((selectedModule) => (
                <div key={selectedModule._id} style={{backgroundColor:'#89C649', textAlign:'center', display: 'inline-block', borderRadius: '.25em', padding: '0.4em 0.7em', marginRight: '10px', marginTop: '45px'}}>
                    <label style={{
                        fontFamily:'"Varela Round",sans-serif',
                        verticalAlign:'middle',
                        color:'#F9F9F9',
                        fontSize:'14px',
                        flex: 1}}
                    >
                        {selectedModule.displayName}
                    </label>
                    <IconButton key={selectedModule._id} onClick={() => this.removeModuleFromSelectedModules(selectedModule._id)} data-tip="Remove" style={{width:'24px', height:'24px', verticalAlign:'middle', paddingLeft:'5px'}} className={classes.iconHover} >
                        <Icon style={{fontSize: '20px'}}>clear</Icon>
                    </IconButton>
                    <ReactTooltip place="bottom" />
                </div>
            ));
        }
        return (
            <div>
                <div style={{textAlign: 'center'}}>
                    <label className={classes.releaseSelectionLabel}>Create a</label>
                    <BuildTypeDropdown onBuildTypeSelection={this.onBuildTypeSelection} />
                    <label className={classes.releaseSelectionLabel}>Application Package from Release:</label>
                    <ReleaseSelectionDropdown onReleaseSelection={this.onReleaseSelection}/>
                </div>
                <div style={{marginBottom:'100px', marginTop:'30px'}}>
                    <Grid>
                        <Row>
                            <Column width="1/2">
                                <h2 className={classes.h2}>
                                    Application Initializer Parameters
                                </h2>
                            </Column>
                            <Column width="1/2">
                                <h2 className={classes.h2}>
                                    Modules
                                </h2>
                            </Column>
                        </Row>
                        <Row>
                            <Column width="1/2">
                                <Row>
                                    <FormControl style={{width: '80%', marginTop: '20px', paddingLeft: '10px'}} required={true} error={this.state.build.utilityId === ""} >
                                        <InputLabel htmlFor="name-simple" className={classes.inputLabel}>Utility
                                            Id</InputLabel>
                                        <Input id="name-simple" name="utilityId" value={this.state.build.utilityId} style={{marginTop: '20px'}} placeholder="i.e. PHI" onChange={this.validateRequiredField} />
                                    </FormControl>
                                </Row>
                                <Row>
                                    <FormControl className={classes.formControl}>
                                        <InputLabel htmlFor="name-simple" className={classes.inputLabel}>Master Model
                                            Id</InputLabel>
                                        <Input id="name-simple" name="masterModelId" value={this.state.build.masterModelId} style={{color: '#333'}} placeholder="i.e. PHI" onChange={this.validateRequiredField}/>
                                    </FormControl>
                                </Row>
                                <Row>
                                    <FormControl className={classes.formControl}>
                                        <InputLabel htmlFor="name-simple" className={classes.inputLabel}>Project
                                            Name</InputLabel>
                                        <Input id="name-simple" name="projectName" value={this.state.build.projectName} style={{color: '#333'}} onChange={this.validateRequiredField}/>
                                    </FormControl>
                                </Row>
                                <Row>
                                    <FormControl className={classes.formControl}>
                                        <InputLabel htmlFor="name-simple"
                                                    className={classes.inputLabel}>Description</InputLabel>
                                        <Input id="name-simple" name="description" value={this.state.build.description} style={{color: '#333'}} onChange={this.validateRequiredField}/>
                                    </FormControl>
                                </Row>
                                <Row>
                                    <FormControl className={classes.formControl}>
                                        <InputLabel htmlFor="name-simple"
                                                    className={classes.inputLabel}>Timezone</InputLabel>
                                        <Input id="name-simple" name="timezone" value={this.state.build.timezone} placeholder="i.e. America/New_York"
                                               style={{color: '#333'}} onChange={this.validateRequiredField}/>
                                    </FormControl>
                                </Row>
                                <Row>
                                    <FormControl className={classes.formControl}>
                                        <InputLabel htmlFor="name-simple" className={classes.inputLabel}>Default
                                            Latitude</InputLabel>
                                        <Input id="name-simple" name="defaultLatitude" value={this.state.build.defaultLatitude} placeholder="i.e. 39.025585" style={{color: '#333'}} onChange={this.validateRequiredField}/>
                                    </FormControl>
                                </Row>
                                <Row>
                                    <FormControl className={classes.formControl}>
                                        <InputLabel htmlFor="name-simple" className={classes.inputLabel}>Default
                                            Longitude</InputLabel>
                                        <Input id="name-simple" name="defaultLongitude" value={this.state.build.defaultLongitude} placeholder="i.e. 77.073898" style={{color: '#333'}} onChange={this.validateRequiredField}/>
                                    </FormControl>
                                </Row>
                            </Column>
                            <Column width="1/2">
                                <Row>
                                    <IntegrationAutosuggest onModuleSelection={this.onModuleSelection}/>
                                </Row>
                                <Row>
                                    {selectedModules}
                                </Row>
                            </Column>
                        </Row>
                    </Grid>
                </div>
                <div>
                    <BuildCreationFormFooter onCancel={this.cancelBuild} onCreate={this.createBuild} buildFormComplete={this.props.buildFormState}/>
                </div>
            </div>
        )
    }
}

BuildCreationForm.propTypes = {
    classes: PropTypes.object.isRequired
};

function mapStateToProps(state) {
    return {
        releases: state.releases,
        modules: state.modules,
        buildFormState: state.buildFormState
    };
}

function matchDispatchToProps(dispatch) {
    return bindActionCreators({
        loadReleases: loadReleases,
        loadModulesForRelease: loadModulesForRelease,
        updateBuildFormState: updateBuildFormState,
        createBuild: createBuild,
        selectView: selectView
    }, dispatch);
}

export default withStyles(styles)(connect(mapStateToProps, matchDispatchToProps)(BuildCreationForm));
