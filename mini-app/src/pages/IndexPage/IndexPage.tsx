import type { FC } from 'react';
import { useState } from 'react';


import {FooterNav} from "@/components/FooterNav/FooterNav.tsx";
import {pages} from "@/navigation/routes.tsx";

export const IndexPage: FC = () => {
    const [page, setPage] = useState(0);
    const { Component } = pages[page];
    return (
        <div style={{ minHeight: '100vh', background: '#212736', paddingBottom: 64 }}>
            <Component />
            <FooterNav
                value={page}
                onChange={setPage}
                items={pages.map(({ label, icon }) => ({ label, icon }))}
            />
        </div>
    );
};