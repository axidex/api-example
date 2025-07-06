import type { FC, ReactNode } from "react";
import "./FooterNav.css";

interface FooterNavProps {
    value: number;
    onChange: (idx: number) => void;
    items: { icon: ReactNode; label: string }[];
}

export const FooterNav: FC<FooterNavProps> = ({ value, onChange, items }) => (
    <nav className="footer-nav">
        {items.map((item, i) => (
            <div
                key={item.label}
                className={`footer-nav-item${i === value ? " active" : ""}`}
                onClick={() => onChange(i)}
            >
                <div className="footer-nav-icon">{item.icon}</div>
                <span>{item.label}</span>
            </div>
        ))}
    </nav>
);