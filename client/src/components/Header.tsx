import { Container, Nav, Navbar, NavDropdown } from "solid-bootstrap";

export default function Header() {
    return (
        <main>
            <Container fluid id="header">
                <Navbar.Brand>YouBlog</Navbar.Brand>
                <Navbar.Toggle />
                <Navbar.Collapse>
                    <Nav>
                        <NavDropdown title="Placeholder">
                            <NavDropdown.Item>Change Username</NavDropdown.Item>
                            <NavDropdown.Item>Change Username</NavDropdown.Item>
                        </NavDropdown>
                    </Nav>
                </Navbar.Collapse>
            </Container>
        </main>
    );
}