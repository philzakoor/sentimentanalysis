import React from 'react';

function WithResultsLoading(Component) {
    return function WihLoadingComponent({ isLoading, ...props }) {
        if (!isLoading) return <Component {...props} />;
        return (
            <p>
                Please Stand By..
            </p>
        );
    };
}

export default WithResultsLoading;
