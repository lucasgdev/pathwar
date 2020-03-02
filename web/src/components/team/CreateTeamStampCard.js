import React, { useState, useEffect } from "react"
import { Link } from "gatsby"
import { isEmpty } from "ramda"
import { StampCard, Form, Button } from "tabler-react"

const CreateTeamStampCard = ({ activeSeason, createTeam }) => {
  const [isFormOpen, setFormOpen] = useState(false)
  const [name, setName] = useState("")
  const [error, setError] = useState(false)

  useEffect(() => {
    if (!isEmpty(name)) {
      setError(false)
    }
  }, [name])

  const handleChange = event => setName(event.target.value)
  const handleFormOpen = event => {
    event.preventDefault()
    setFormOpen(!isFormOpen)
  }

  const submitTeamCreate = async event => {
    event.preventDefault()
    if (isEmpty(name)) {
      setError(true)
      return
    } else {
      await createTeam(activeSeason.id, name)
      setFormOpen(false)
    }
  }

  return (
    <>
      <StampCard
        color="blue"
        icon="users"
        header={
          <Link to="/" onClick={handleFormOpen}>
            <small>Create new team</small>
          </Link>
        }
        footer={"Ahoy! Create a new team"}
      />
      {isFormOpen && (
        <form onSubmit={submitTeamCreate}>
          <Form.FieldSet>
            <Form.Group isRequired label="Name">
              <Form.Input
                name="name"
                onChange={handleChange}
                invalid={error}
                cross={error}
                feedback={error && "Please, insert a name"}
              />
            </Form.Group>
            <Form.Group>
              <Button type="submit" color="primary" className="ml-auto">
                Send
              </Button>
            </Form.Group>
          </Form.FieldSet>
        </form>
      )}
    </>
  )
}

export default CreateTeamStampCard