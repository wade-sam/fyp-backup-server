import React from 'react';
import CheckboxTree from 'react-checkbox-tree';
 
const nodes = [{
    value: 'mars',
    label: 'Mars',
    children: [
        { value: 'phobos', label: 'Phobos' },
        { value: 'deimos', label: 'Deimos' },
    ],
}];
 
class Widget extends React.Component {
    state = {
        checked: [],
        expanded: [],
    };

    
 
    render() {
        console.log(this.state.checked);
        return (
            <CheckboxTree
                nodes={nodes}
                checked={this.state.checked}
                expanded={this.state.expanded}
                onCheck={checked => this.setState({ checked })}
                onExpand={expanded => this.setState({ expanded })}
                nativeCheckboxes={true}
            />
        );
    }
}
export default Widget