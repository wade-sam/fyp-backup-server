import React from 'react';
import CheckboxTree from 'react-checkbox-tree';

class DirectoryScan extends React.Component{
    state = {
        checked: [],
        expanded: [],
        data: this.props.scan
    };
    
    
    render(){
        const data = this.props.scan
        console.log("DirectoryScan reached, ", this.props.scan);
        return(
            <React.Fragment>
                <CheckboxTree
                    nodes={data}
                    key={this.state.data.value}
                    checked={this.state.checked}
                    expanded={this.state.expanded}
                    onCheck={checked => this.setState({ checked })}
                    onExpand={expanded => this.setState({ expanded })}
                />
            </React.Fragment>
        );
    }
}
export default DirectoryScan