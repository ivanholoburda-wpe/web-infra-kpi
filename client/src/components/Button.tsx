import React from 'react';

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    variant: 'add' | 'update' | 'edit' | 'delete';
}

const Button: React.FC<ButtonProps> = ({variant, children, ...props}) => {
    return (
        <button className={`button button--${variant}`} {...props}>
            {children}
        </button>
    );
};

export default Button;