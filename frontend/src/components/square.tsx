import React from 'react';

interface State {}
interface Props {
  value: string;
  onClick: () => void;
}

export class Square extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {};
  }

  render() {
    return (
      <button className="square" onClick={this.props.onClick}>
        {this.props.value}
      </button>
    );
  }
}
