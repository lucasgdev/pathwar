import {
  GET_ALL_SEASONS_SUCCESS,
  GET_ALL_SEASONS_FAILED,
  GET_ALL_SEASON_TEAMS_SUCCESS,
  GET_ALL_SEASON_TEAMS_FAILED,
  SET_DEFAULT_SEASON,
  SET_ACTIVE_SEASON,
  SET_ACTIVE_SEASON_FAILED,
  FETCH_PREFERENCES_SUCCESS,
  FETCH_PREFERENCES_FAILED,
  SET_CHALLENGES_LIST,
  SET_CHALLENGES_LIST_FAILED,
  GET_CHALLENGE_DETAILS_SUCCESS,
  GET_CHALLENGE_DETAILS_FAILED,
  GET_TEAM_DETAILS_SUCCESS,
  GET_TEAM_DETAILS_FAILED
} from "../constants/actionTypes"

import {
  getAllSeasons,
  getAllSeasonTeams,
  postPreferences,
  getChallenges,
  getChallengeDetails,
  getTeamDetails
} from "../api/seasons"

import { fetchUserSession as fetchUserSessionAction } from "./userSession";

export const fetchPreferences = (seasonID) => async dispatch => {
  try {
    await postPreferences(seasonID)

    dispatch({
      type: FETCH_PREFERENCES_SUCCESS
    });

    dispatch(fetchUserSessionAction(false));

  } catch(error) {
    dispatch({
      type: FETCH_PREFERENCES_FAILED,
      payload: { error }
    });
  }
}

export const setActiveSeason = (seasonData) => async dispatch => {
  try {

      dispatch({
        type: SET_ACTIVE_SEASON,
        payload: { activeSeason: seasonData }
      });
  }
  catch(error) {
    dispatch({ type: SET_ACTIVE_SEASON_FAILED, payload: { error }});
    alert("Set season active failed, please try again!")
  }
}

export const setDefaultSeason = (seasonData) => async dispatch => {
  dispatch({
    type: SET_DEFAULT_SEASON,
    payload: { defaultSeason: seasonData }
  });
}

export const fetchAllSeasonTeams = (seasonID) => async dispatch => {
  try {
    const response = await getAllSeasonTeams(seasonID);
    const allTeams = response.data.items;

    dispatch({
      type: GET_ALL_SEASON_TEAMS_SUCCESS,
      payload: { allTeams: allTeams }
    })
  } catch (error) {
    dispatch({ type: GET_ALL_SEASON_TEAMS_FAILED, payload: { error } });
  }
}

export const fetchTeamDetails = (teamID) => async dispatch => {
  try {
    const response = await getTeamDetails(teamID);
    const detailsResponse = response.data.item;

    dispatch({
      type: GET_TEAM_DETAILS_SUCCESS,
      payload: {
        team: detailsResponse,
      }
    })

  } catch (error) {
    dispatch({
      type: GET_TEAM_DETAILS_FAILED,
      payload: { error }
    })
  }
}

export const fetchAllSeasons = () => async dispatch => {
  try {
    const response = await getAllSeasons();
    const allSeasons = response.data.items;

    dispatch({
      type: GET_ALL_SEASONS_SUCCESS,
      payload: { allSeasons: allSeasons }
    })
  } catch (error) {
    dispatch({ type: GET_ALL_SEASONS_FAILED, payload: { error } });
  }
}

export const fetchChallengeDetail = (challengeID) => async dispatch => {
  try {
    const response = await getChallengeDetails(challengeID);

    dispatch({
      type: GET_CHALLENGE_DETAILS_SUCCESS,
      payload: { challenge: response.data.item }
    });

  } catch (error) {
    dispatch({ type: GET_CHALLENGE_DETAILS_FAILED, payload: { error } });
  }
};

export const fetchChallenges = (seasonID) => async dispatch => {
  try {
    const response = await getChallenges(seasonID);
    dispatch({
      type: SET_CHALLENGES_LIST,
      payload: { activeChallenges: response.data.items }
    });
  } catch (error) {
    dispatch({ type: SET_CHALLENGES_LIST_FAILED, payload: { error } });
  }
};
