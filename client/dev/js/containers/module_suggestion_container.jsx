import React, {Component} from "react";
import PropTypes from "prop-types";
import Autosuggest from "react-autosuggest";
import Paper from "material-ui/Paper";
import {MenuItem} from "material-ui/Menu";
import match from "autosuggest-highlight/match";
import parse from "autosuggest-highlight/parse";
import {withStyles} from "material-ui/styles";
import {connect} from "react-redux";
import Input from "material-ui/Input";
import InputLabel from "material-ui/Input/InputLabel";
import FormControl from "material-ui/Form/FormControl";

function renderInput(inputProps) {
    const { classes, home, value, ref, ...other} = inputProps;
    return (
        <FormControl style={{width: '80%', marginTop: '20px', paddingLeft: '10px'}}>
            <InputLabel htmlFor="name-simple" style={{paddingLeft:'10px'}}>Search for Modules</InputLabel>
            <Input
                id="name-simple"
                style={{marginTop: '20px'}}
                autoFocus={home}
                inputRef={ref}
                value={value}
                {...other}
            />
        </FormControl>
    );
}

function renderSuggestion(suggestion, { query, isHighlighted }) {
    const matches = match(suggestion.displayName, query);
    const parts = parse(suggestion.displayName, matches);

    return (
        <MenuItem selected={isHighlighted} component="div">
            <div>
                {parts.map((part, index) => {
                    return part.highlight
                        ? <span key={index} style={{ fontWeight: 300 }}>
                {part.text}
              </span>
                        : <strong key={index} style={{ fontWeight: 500 }}>
                            {part.text}
                        </strong>;
                })}
            </div>
        </MenuItem>
    );
}

function renderSuggestionsContainer(options) {
    const { containerProps, children } = options;

    return (
        <Paper {...containerProps} square>
            {children}
        </Paper>
    );
}

function getSuggestions(value, modules) {
    const inputValue = value.trim().toLowerCase();
    const inputLength = inputValue.length;
    let count = 0;

    return inputLength === 0
        ? []
        : modules.filter(suggestion => {
            const keep =
                count < 5 && checkUIModule(suggestion) && suggestion.displayName.toLowerCase().slice(0, inputLength) === inputValue;
            if (keep) {
                count += 1;
            }
            return keep;
        });
}

function checkUIModule(module) {
    return !module.buildModule && !module.builderDependencies && !module.buildInfrastructure
}

const styles = theme => ({
    container: {
        flexGrow: 1,
        position: 'relative',
    },
    suggestionsContainerOpen: {
        position: 'absolute',
        marginTop: theme.spacing.unit,
        marginBottom: theme.spacing.unit * 3,
        left: 0,
        right: 0,
        zIndex:'10000'
    },
    suggestion: {
        display: 'block',
    },
    suggestionsList: {
        margin: 0,
        padding: 0,
        listStyleType: 'none'
    },
    textField: {
        width: '100%',
    },
});

class IntegrationAutosuggest extends Component {
    constructor(props) {
        super(props);
        this.state = {
            value: '',
            suggestions: []
        };
        this.handleSuggestionsFetchRequested = this.handleSuggestionsFetchRequested.bind(this);
        this.handleSuggestionsClearRequested = this.handleSuggestionsClearRequested.bind(this);
        this.handleChange = this.handleChange.bind(this);
        this.getSuggestionValue = this.getSuggestionValue.bind(this);
        this.buildModuleNameString = this.buildModuleNameString.bind(this);
    }

    handleSuggestionsFetchRequested({ value }) {
        this.setState({
            suggestions: getSuggestions(value, this.props.modules),
        });
    };

    handleSuggestionsClearRequested() {
        this.setState({
            suggestions: [],
            value: "",
        });
    };

    handleChange(event, { newValue }){
        this.setState({
            value: newValue,
        });
    };

    getSuggestionValue(suggestion) {
        this.props.onModuleSelection(suggestion);
        return suggestion.displayName;
    }

    buildModuleNameString() {
        let moduleNames = this.props.modules.filter(module => {
            return checkUIModule(module) //&& this.state.selectedModules.indexOf(module) <= -1
        }).map(module => {
            return module.displayName
        }).join(",  ");
        moduleNames += "...";
        return moduleNames
    }

    render() {
        const {classes} = this.props;

        if (!this.props.modules) {
            return null
        }
        return (
            <Autosuggest
                theme={{
                    container: classes.container,
                    suggestionsContainerOpen: classes.suggestionsContainerOpen,
                    suggestionsList: classes.suggestionsList,
                    suggestion: classes.suggestion,
                }}
                renderInputComponent={renderInput}
                suggestions={this.state.suggestions}
                onSuggestionsFetchRequested={this.handleSuggestionsFetchRequested}
                onSuggestionsClearRequested={this.handleSuggestionsClearRequested}
                renderSuggestionsContainer={renderSuggestionsContainer}
                getSuggestionValue={this.getSuggestionValue}
                renderSuggestion={renderSuggestion}
                inputProps={{
                    autoFocus: true,
                    classes,
                    placeholder: this.buildModuleNameString(),
                    value: this.state.value,
                    onChange: this.handleChange,
                }}
            />
        );
    }
}

IntegrationAutosuggest.propTypes = {
    classes: PropTypes.object.isRequired,
    onModuleSelection: PropTypes.func.isRequired
};

function mapStateToProps(state) {
    return {
        modules: state.modules
    };
}

export default withStyles(styles)(connect(mapStateToProps)(IntegrationAutosuggest));
