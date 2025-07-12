import type {ComponentType, ReactNode} from 'react';

import { InitDataPage } from '@/pages/InitDataPage.tsx';
import { LaunchParamsPage } from '@/pages/LaunchParamsPage.tsx';
import { ThemeParamsPage } from '@/pages/ThemeParamsPage.tsx';
import { MainPage } from '@/pages/MainPage/MainPage.tsx';
import {MdHistory, MdHome, MdLeaderboard, MdPerson} from "react-icons/md";
import {IndexPage} from "@/pages/IndexPage/IndexPage.tsx";

interface Route {
  path: string;
  Component: ComponentType;
}

export const routes: Route[] = [
  { path: "/", Component: IndexPage },
];

interface Page {
  label: string;
  Component: ComponentType;
  icon: ReactNode
}

export const pages: Page[] = [
  { label: "Main", Component: MainPage, icon: <MdHome />},
  { label: "Init", Component: InitDataPage, icon: <MdLeaderboard /> },
  { label: "LaunchParams", Component: LaunchParamsPage, icon: <MdPerson /> },
  { label: "ThemeParams", Component: ThemeParamsPage, icon: <MdHistory /> },
];
