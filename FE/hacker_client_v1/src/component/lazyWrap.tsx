import { FC,LazyExoticComponent,Suspense } from 'react';

interface LazyWrapProps {
    Component: LazyExoticComponent<FC>;
}

const LazyWrap:FC<LazyWrapProps> = ({Component})=>{
    return (
        <Suspense fallback={<div>loading...</div>}>
            <Component />
        </Suspense>
    );
};

export default LazyWrap;